import { TransactionCategory } from "@prisma/client";

export type CreateAccountResponse = {
  id: string;
};

export type UpdateAccountResponse = {
  id: string;
};

export type AccountTransaction = {
  id: string;
  amount: number;
  isExpense: boolean;
  category: TransactionCategory;
  occurredAt: Date;
};

export type AccountOverview = {
  id: string;
  name: string;
  createdAt: Date;
};

export type AccountDetails = {
  id: string;
  name: string;
  transactions: AccountTransaction[];
};

export type GetAccountsResponse = AccountOverview[];

export type GetAccountDetailsResponse = AccountDetails;
