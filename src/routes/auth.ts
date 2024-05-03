import { jwtAccessMiddleware, jwtRefreshMiddleware } from "@/middlewares/jwt";
import { AuthSignInSchema, AuthSignUpSchema } from "@/schemas/auth";
import { AppEnv } from "@/schemas/env";
import { AuthJwtPayload } from "@/types/auth";
import {
  refreshTokenUsecase,
  signInUsecase,
  signUpUsecase,
} from "@/usecase/auth";
import { zValidator } from "@hono/zod-validator";
import { Hono } from "hono";
import { env } from "hono/adapter";
import httpStatus = require("http-status");

export const authRoutes = new Hono();

authRoutes.post("/signup", zValidator("json", AuthSignUpSchema), async (c) => {
  const dto = c.req.valid("json");

  const result = await signUpUsecase(dto);

  return c.json(result, httpStatus.CREATED);
});

authRoutes.post("/signin", zValidator("json", AuthSignInSchema), async (c) => {
  const dto = c.req.valid("json");
  const { JWT_SECRET, JWT_REFRESH_SECRET } = env<AppEnv>(c);

  const result = await signInUsecase(dto, {
    access: JWT_SECRET,
    refresh: JWT_REFRESH_SECRET,
  });

  return c.json(result, httpStatus.OK);
});

authRoutes.post("/tokens/refresh", jwtRefreshMiddleware, async (c) => {
  const { JWT_SECRET, JWT_REFRESH_SECRET } = env<AppEnv>(c);

  const payload: AuthJwtPayload = c.get("jwtPayload");

  const result = await refreshTokenUsecase(payload, {
    access: JWT_SECRET,
    refresh: JWT_REFRESH_SECRET,
  });

  return c.json(result, httpStatus.OK);
});

authRoutes.get("/me", jwtAccessMiddleware, async (c) => {
  const payload: AuthJwtPayload = c.get("jwtPayload");

  console.log(payload);

  return c.body(null, httpStatus.OK);
});
