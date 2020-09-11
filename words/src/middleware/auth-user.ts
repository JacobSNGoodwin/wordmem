import * as fs from "fs";
import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";

interface TokenClaims {
  user: UserClaims;
}

interface UserClaims {
  uid: string;
  email: string;
  name: string;
  imageUrl: string;
  website: string;
}

// For appending a current user to Express Request
declare global {
  namespace Express {
    interface Request {
      currentUser?: UserClaims;
    }
  }
}

// maybe this needs to be injected
const pubKey = fs.readFileSync("src/rsa_public.pem");

export const authUser = (req: Request, res: Response, next: NextFunction) => {
  const authHeader = req.header("Authorization");

  const token = authHeader?.split(" ")[1];

  if (!token) {
    const err = new Error("No token included");
    return next(err);
  }

  try {
    // can do the following sincec jwt.verify throws
    // for invalid or unverified token
    const tokenPayload = jwt.verify(token, pubKey) as TokenClaims;
    req.currentUser = tokenPayload.user;
    console.log(req.currentUser);
  } catch (e) {
    // TODO - Throw an authorization error to be handled by middleware
    console.log("Error verifying token");
  }

  next();
};
