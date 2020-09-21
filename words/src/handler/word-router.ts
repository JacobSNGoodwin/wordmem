import express, { Request, Response, Router, NextFunction } from "express";
import { body, query } from "express-validator";

import { requireAuth } from "../middleware/require-auth";
import { serviceContainer } from "../injection";
import { validateRequest } from "../middleware/validate-request";
import { BadRequestError } from "../errors/bad-request-error";

export const createWordRouter = (): Router => {
  const wordRouter = express.Router();
  const { wordService } = serviceContainer.services;

  wordRouter.use(requireAuth);

  wordRouter.get(
    "/",
    [
      query("fib")
        .optional({ nullable: true })
        .isBoolean()
        .withMessage("Must be null or a boolean"),
      query("page")
        .optional({ nullable: true })
        .isInt()
        .withMessage("Must be null or an integer"),
      query("limit")
        .optional({ nullable: true })
        .isInt({ max: 100 })
        .withMessage("Must be null or an integer less than or equal to 100"),
    ],
    validateRequest,
    async (req: Request, res: Response, next: NextFunction) => {
      const limit = parseInt(req.query["limit"] as string);
      const page = parseInt(req.query["page"] as string);
      const isFibo = req.query["fib"] === "true" ? true : false;

      try {
        const wordList = await wordService.getWords({
          userId: req.currentUser!.uid,
          limit,
          page,
          isFibo,
        });

        res.status(200).json(wordList);
      } catch (err) {
        next(err);
      }
    }
  );

  wordRouter.post(
    "/",
    [
      body("word").notEmpty().trim().withMessage("required"),
      body("definition").notEmpty().trim().withMessage("required"),
      body("refUrl").optional().isURL().trim().withMessage("url"),
      body("emailReminder").optional().isBoolean().withMessage("boolean"),
    ],
    validateRequest,
    async (req: Request, res: Response, next: NextFunction) => {
      const { word, refUrl, emailReminder, definition } = req.body;

      try {
        const created = await wordService.addWord({
          word,
          definition,
          refUrl,
          emailReminder,
          email: req.currentUser!.email,
          uid: req.currentUser!.uid,
        });

        res.status(201).json(created);
      } catch (err) {
        next(err);
      }
    }
  );

  // using a post request with list of posts
  // not sure if this is totally RESTful, but what the hell is?
  // https://stackoverflow.com/questions/21863326/delete-multiple-records-using-rest/30933909
  wordRouter.post(
    "/delete",
    [
      body("wordIds")
        .isArray({ min: 1 })
        .withMessage("must be array with non-zero length"),
      body("wordIds.*").isUUID().withMessage("array must contain UUIDs"),
    ],
    validateRequest,
    async (req: Request, res: Response, next: NextFunction) => {
      try {
        const deletedIds = await wordService.deleteWords(req.body.wordIds);
        return res.status(200).json(deletedIds);
      } catch (err) {
        next(err);
      }
    }
  );

  wordRouter.put(
    "/:id",
    [
      body("word")
        .isString()
        .isLength({ min: 1 })
        .trim()
        .withMessage("required"),
      body("definition")
        .exists({ checkNull: true })
        .isString()
        .trim()
        .withMessage("required"),
      body("refUrl")
        .exists({ checkNull: true })
        .if(body("refUrl").notEmpty())
        .isURL()
        .trim()
        .withMessage("url"),
      body("emailReminder")
        .exists({ checkNull: true })
        .isBoolean()
        .withMessage("boolean"),
      body("startDate")
        .exists({ checkNull: true })
        .notEmpty()
        .withMessage("date"),
    ],
    validateRequest,
    async (req: Request, res: Response, next: NextFunction) => {
      const {
        word,
        refUrl,
        emailReminder,
        definition,
        startDate: strDate,
      } = req.body;

      // parse date
      let startDate: Date;
      try {
        startDate = new Date(strDate as string);
      } catch (err) {
        console.error("Invalid date string!", err);
        throw new BadRequestError(
          "startDate must be a string that can be parsed as a date"
        );
      }

      try {
        const updated = await wordService.updateWord({
          id: req.params.id,
          uid: req.currentUser!.uid,
          word,
          definition,
          refUrl,
          emailReminder,
          startDate,
        });

        res.status(200).json(updated);
      } catch (err) {
        next(err);
      }
    }
  );

  return wordRouter;
};
