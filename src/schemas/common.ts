import { z } from "zod";

export const PaginationSchema = z
  .object({
    cursor: z.string(),
    limit: z.number().int().positive(),
  })
  .partial();

export type PaginationDTO = z.infer<typeof PaginationSchema>;
