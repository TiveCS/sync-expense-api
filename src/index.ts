import { Hono } from "hono";
import { logger } from "hono/logger";
import { authRoutes } from "./routes/auth";
import { errorFilters } from "./exceptions/filters";

const app = new Hono();

app.use(logger()).onError(errorFilters).route("/auth", authRoutes);

export default {
  port: process.env.PORT || 8080,
  fetch: app.fetch,
};
