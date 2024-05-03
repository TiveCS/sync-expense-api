import { jwtAccessMiddleware } from "@/middlewares/jwt";
import { CreateAccountSchema, UpdateAccountSchema } from "@/schemas/accounts";
import { AuthJwtPayload } from "@/types/auth";
import {
  createAccountUsecase,
  deleteAccountUsecase,
  getAccountDetailsUsecase,
  getAccountsUsecase,
  updateAccountUsecase,
} from "@/usecase/accounts";
import { zValidator } from "@hono/zod-validator";
import { Hono } from "hono";
import httpStatus from "http-status";

export const accountsRoutes = new Hono().use(jwtAccessMiddleware);

accountsRoutes.get("/", async (c) => {
  const userPayload: AuthJwtPayload = c.get("jwtPayload");

  const result = await getAccountsUsecase(userPayload.sub);

  return c.json(result, { status: httpStatus.OK });
});

accountsRoutes.post("/", zValidator("json", CreateAccountSchema), async (c) => {
  const dto = c.req.valid("json");
  const userPayload: AuthJwtPayload = c.get("jwtPayload");

  const result = await createAccountUsecase(userPayload.sub, dto);

  return c.json(result, { status: httpStatus.CREATED });
});

accountsRoutes.get("/accounts/:id", async (c) => {
  const accountId: string = c.req.param("id");

  const result = await getAccountDetailsUsecase(accountId);

  return c.json(result, { status: httpStatus.OK });
});

accountsRoutes.put(
  "/:id",
  zValidator("json", UpdateAccountSchema),
  async (c) => {
    const accountId: string = c.req.param("id");
    const dto = c.req.valid("json");

    const result = await updateAccountUsecase(accountId, dto);

    return c.json(result, { status: httpStatus.OK });
  }
);

accountsRoutes.delete("/:id", async (c) => {
  const accountId: string = c.req.param("id");

  await deleteAccountUsecase(accountId);

  return c.json(null, { status: httpStatus.NO_CONTENT });
});
