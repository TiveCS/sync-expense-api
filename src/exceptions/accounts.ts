import { AppExceptionProperties } from "@/types/exception";
import httpStatus from "http-status";

export const AccountsExceptions = {
  ACCOUNT_NOT_FOUND: {
    code: "ACCOUNTS_EXCEPTION.ACCOUNT_NOT_FOUND",
    message: "Account not found",
    status: httpStatus.NOT_FOUND,
  },
} satisfies Record<string, AppExceptionProperties>;
