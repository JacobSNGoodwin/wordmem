import express, { Request, Response } from "express";
import { body, validationResult } from "express-validator";

import { requireAuth } from "../middleware/require-auth";
import { RequestValidationError } from "../errors/request-validation-error";
import { WordService } from "../service/word-service";

const wordRouter = express.Router();

wordRouter.use(requireAuth);

wordRouter.get("/", (req: Request, res: Response) => {
  res.json({
    user: req.currentUser,
    reqBody: req.body,
  });
});

wordRouter.post(
  "/",
  [
    body("word").not().isEmpty().trim().withMessage("required"),
    body("refUrl").optional().isURL().trim().withMessage("url"),
    body("emailReminder").optional().isBoolean().withMessage("boolean"),
  ],
  (req: Request, res: Response) => {
    const errors = validationResult(req);

    if (!errors.isEmpty()) {
      throw new RequestValidationError(errors.array());
    }

    res.status(200).json({
      word: req.body.word,
      refUrl: req.body.refUrl,
      emailReminder: req.body.emailReminder,
    });
  }
);

export { wordRouter };
