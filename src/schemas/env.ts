import { z } from "zod";

export const AppEnvSchema = z.object({
  DATABASE_URL: z.string(),
  PORT: z.coerce.number(),
  JWT_SECRET: z.string(),
  JWT_REFRESH_SECRET: z.string(),
});

export type AppEnv = z.infer<typeof AppEnvSchema>;
