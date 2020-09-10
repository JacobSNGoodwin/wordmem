import express from "express";

const appRouter = express.Router();

appRouter.get("/", (req, res) => {
  res.json(req.body);
});

export default appRouter;
