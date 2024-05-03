import { AppExceptionProperties } from "@/types/exception";
import httpStatus from "http-status";

export const AuthExceptions = {
  USER_EXISTS: {
    code: "AUTH_EXCEPTIONS.USER_EXISTS",
    message: "User already exists",
    status: httpStatus.BAD_REQUEST,
  },
  INVALID_CREDENTIALS: {
    code: "AUTH_EXCEPTIONS.INVALID_CREDENTIALS",
    message: "Invalid credentials",
    status: httpStatus.UNAUTHORIZED,
  },
  INVALID_TOKEN: {
    code: "AUTH_EXCEPTIONS.INVALID_TOKEN",
    message: "Invalid token",
    status: httpStatus.UNAUTHORIZED,
  },
} satisfies Record<string, AppExceptionProperties>;
