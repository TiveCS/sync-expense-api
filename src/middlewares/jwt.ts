import { AppEnv } from "@/schemas/env";
import { env } from "hono/adapter";
import { createMiddleware } from "hono/factory";
import { jwt } from "hono/jwt";

export const jwtAccessMiddleware = createMiddleware(async (c, next) => {
  const { JWT_SECRET } = env<AppEnv>(c);

  const jwtMiddleware = jwt({
    secret: JWT_SECRET,
  });

  return jwtMiddleware(c, next);
});

export const jwtRefreshMiddleware = createMiddleware(async (c, next) => {
  const { JWT_REFRESH_SECRET } = env<AppEnv>(c);

  const jwtRefreshMiddleware = jwt({
    secret: JWT_REFRESH_SECRET,
  });

  return jwtRefreshMiddleware(c, next);
});
