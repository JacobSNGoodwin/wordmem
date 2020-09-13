import express, { Request, Response, Router, NextFunction } from "express";
import { body, check } from "express-validator";

import { requireAuth } from "../middleware/require-auth";
import { serviceContainer } from "../injection";
import { validateRequest } from "../middleware/validate-request";

export const createWordRouter = (): Router => {
  const wordRouter = express.Router();
  const { wordService } = serviceContainer.services;

  wordRouter.use(requireAuth);

  wordRouter.get(
    "/",
    async (req: Request, res: Response, next: NextFunction) => {
      try {
        const wordList = await wordService.getWords(req.currentUser!.uid);

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
        const created = await wordService.addWord(
          { word, definition, refUrl, emailReminder },
          { email: req.currentUser!.email, id: req.currentUser!.uid }
        );

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

  return wordRouter;
};
