import { AppExceptionProperties } from "@/types/exception";
import httpStatus from "http-status";

export const TransactionExceptions = {
  NOT_FOUND: {
    code: "TRANSACTION_EXCEPTION.NOT_FOUND",
    message: "Transaction not found",
    status: httpStatus.NOT_FOUND,
  },
} satisfies Record<string, AppExceptionProperties>;
