import { jwtAccessMiddleware } from "@/middlewares/jwt";
import { PaginationSchema } from "@/schemas/common";
import {
  CreateTransactionSchema,
  UpdateTransactionSchema,
} from "@/schemas/transactions";
import {
  createTransactionUsecase,
  deleteTransactionUsecase,
  getAccountTransactionsUsecase,
  updateTransactionUsecase,
} from "@/usecase/transactions";
import { zValidator } from "@hono/zod-validator";
import { Hono } from "hono";
import httpStatus from "http-status";

export const transactionsRoutes = new Hono()
  .basePath("/accounts/:accountId/transactions")
  .use(jwtAccessMiddleware);

transactionsRoutes.post(
  "/",
  zValidator("json", CreateTransactionSchema),
  async (c) => {
    const accountId = c.req.param("accountId");
    const dto = c.req.valid("json");

    const result = await createTransactionUsecase(accountId, dto);

    return c.json(result, { status: httpStatus.CREATED });
  }
);

transactionsRoutes.get(
  "/",
  zValidator("query", PaginationSchema),
  async (c) => {
    const accountId = c.req.param("accountId");
    const pagination = c.req.valid("query");

    const result = await getAccountTransactionsUsecase(accountId, pagination);

    return c.json(result, { status: httpStatus.OK });
  }
);

transactionsRoutes.put(
  "/:id",
  zValidator("json", UpdateTransactionSchema),
  async (c) => {
    const transactionId = c.req.param("id");
    const dto = c.req.valid("json");

    const result = await updateTransactionUsecase(transactionId, dto);

    return c.json(result, { status: httpStatus.OK });
  }
);

transactionsRoutes.delete("/:id", async (c) => {
  const transactionId = c.req.param("id");

  await deleteTransactionUsecase(transactionId);

  return c.json(null, { status: httpStatus.NO_CONTENT });
});
