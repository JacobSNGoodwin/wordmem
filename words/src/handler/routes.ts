import express from "express";

import { requireAuth } from "../middleware/require-auth";

const wordRouter = express.Router();

wordRouter.use(requireAuth);

wordRouter.get("/", (req, res) => {
  res.json({
    user: req.currentUser,
    reqBody: req.body,
  });
});

export { wordRouter };
