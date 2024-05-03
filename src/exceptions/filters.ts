import { ErrorHandler } from "hono";
import httpStatus from "http-status";
import { AppException } from "./common";

export const errorFilters: ErrorHandler = (err, c) => {
  if (err instanceof AppException) {
    return err.getResponse();
  }

  return c.json(
    { message: err.message || "Unhandled Error" },
    { status: httpStatus.INTERNAL_SERVER_ERROR }
  );
};
