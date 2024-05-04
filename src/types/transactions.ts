import { Transaction, TransactionCategory } from "@prisma/client";

export type CreateTransactionResponse = {
  id: string;
  occurredAt: Date;
};

export type UpdateTransactionResponse = {
  id: string;
};

export type AccountTransaction = {
  id: string;
  amount: number;
  isExpense: boolean;
  category: TransactionCategory;
  occurredAt: Date;
};

export type GetAccountTransactionsResponse = AccountTransaction[];
