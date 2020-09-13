import express, { Request, Response, Router, NextFunction } from "express";
import { body } from "express-validator";

import { requireAuth } from "../middleware/require-auth";
import { serviceContainer } from "../injection";
import { validateRequest } from "../middleware/validate-request";
import { nextTick } from "process";

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
      body("word").not().isEmpty().trim().withMessage("required"),
      body("refUrl").optional().isURL().trim().withMessage("url"),
      body("emailReminder").optional().isBoolean().withMessage("boolean"),
    ],
    validateRequest,
    async (req: Request, res: Response, next: NextFunction) => {
      const { word, refUrl, emailReminder } = req.body;

      try {
        const created = await wordService.addWord(
          { word, refUrl, emailReminder },
          { email: req.currentUser!.email, id: req.currentUser!.uid }
        );

        res.status(201).json(created);
      } catch (err) {
        next(err);
      }
    }
  );

  return wordRouter;
};
