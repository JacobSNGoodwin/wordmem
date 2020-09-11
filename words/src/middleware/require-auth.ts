import { Request, Response, NextFunction } from "express";

export const requireAuth = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  if (!req.currentUser) {
    throw new Error("Not authorized!");
  }

  next();
};
