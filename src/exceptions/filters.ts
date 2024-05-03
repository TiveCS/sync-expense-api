import { ErrorHandler } from "hono";
import { HTTPException } from "hono/http-exception";
import httpStatus from "http-status";

export const errorFilters: ErrorHandler = (err, c) => {
  if (err instanceof HTTPException) {
    return err.getResponse();
  }

  return c.json(
    { message: err.message || "Unhandled Error" },
    { status: httpStatus.INTERNAL_SERVER_ERROR }
  );
};
