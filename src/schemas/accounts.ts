import { z } from "zod";

export const CreateAccountSchema = z.object({
  name: z.string().min(1),
});

export const UpdateAccountSchema = z.object({
  name: z.string().min(1),
});

export type CreateAccountDTO = z.infer<typeof CreateAccountSchema>;
export type UpdateAccountDTO = z.infer<typeof UpdateAccountSchema>;
