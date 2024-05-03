import { AppExceptionProperties } from "@/types/exception";
import { HTTPException } from "hono/http-exception";
import { StatusCode } from "hono/utils/http-status";
import httpStatus from "http-status";

export class AppException extends HTTPException {
  private code: string;

  constructor({
    code,
    message,
    status = httpStatus.INTERNAL_SERVER_ERROR,
  }: AppExceptionProperties) {
    super(status as StatusCode, {
      message,
      cause: code,
    });
    this.code = code;
  }

  getResponse(): Response {
    return new Response(
      JSON.stringify({ code: this.code, message: this.message }),
      {
        status: this.status,
        statusText: this.message,
        headers: {
          "Content-Type": "application/json; charset=utf-8",
        },
      }
    );
  }
}
