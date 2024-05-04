import { Hono } from "hono";
import { logger } from "hono/logger";
import { authRoutes } from "./routes/auth";
import { errorFilters } from "./exceptions/filters";
import { accountsRoutes } from "./routes/accounts";
import { transactionsRoutes } from "./routes/transactions";

const app = new Hono();

app
  .use(logger())
  .onError(errorFilters)
  .route("/auth", authRoutes)
  .route("/accounts", accountsRoutes)
  .route("/", transactionsRoutes);

export default {
  port: process.env.PORT || 8080,
  fetch: app.fetch,
};
