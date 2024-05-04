import { TransactionCategory } from "@prisma/client";
import { z } from "zod";

export const CreateTransactionSchema = z.object({
  amount: z.number().positive(),
  category: z.nativeEnum(TransactionCategory),
  isExpense: z.boolean(),
  occurredAt: z.date(),
});

export const UpdateTransactionSchema = z.object({
  amount: z.number().positive().optional(),
  category: z.nativeEnum(TransactionCategory).optional(),
  isExpense: z.boolean().optional(),
  occurredAt: z.date().optional(),
});

export type CreateTransactionDTO = z.infer<typeof CreateTransactionSchema>;
export type UpdateTransactionDTO = z.infer<typeof UpdateTransactionSchema>;
