import { AppException } from "@/exceptions/common";
import { TransactionExceptions } from "@/exceptions/transactions";
import { prisma } from "@/lib/prisma";
import { PaginationDTO } from "@/schemas/common";
import {
  CreateTransactionDTO,
  UpdateTransactionDTO,
} from "@/schemas/transactions";
import {
  CreateTransactionResponse,
  GetAccountTransactionsResponse,
  UpdateTransactionResponse,
} from "@/types/transactions";
import { Prisma } from "@prisma/client";

export async function createTransactionUsecase(
  accountId: string,
  dto: CreateTransactionDTO
): Promise<CreateTransactionResponse> {
  const transaction = await prisma.transaction.create({
    data: {
      amount: dto.amount,
      category: dto.category,
      occurredAt: dto.occurredAt,
      isExpense: dto.isExpense,
      account: { connect: { id: accountId } },
    },
    select: { id: true, occurredAt: true },
  });

  return {
    id: transaction.id,
    occurredAt: transaction.occurredAt,
  };
}

export async function updateTransactionUsecase(
  transactionId: string,
  dto: UpdateTransactionDTO
): Promise<UpdateTransactionResponse> {
  try {
    const transaction = await prisma.transaction.update({
      where: { id: transactionId },
      data: {
        amount: dto.amount,
        category: dto.category,
        occurredAt: dto.occurredAt,
        isExpense: dto.isExpense,
      },
    });

    return {
      id: transaction.id,
    };
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      if (error.code === "P2025") {
        throw new AppException(TransactionExceptions.NOT_FOUND);
      }
    }
    throw error;
  }
}

export async function deleteTransactionUsecase(
  transactionId: string
): Promise<void> {
  try {
    await prisma.transaction.delete({
      where: { id: transactionId },
    });
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      if (error.code === "P2025") {
        throw new AppException(TransactionExceptions.NOT_FOUND);
      }
    }
    throw error;
  }
}

export async function getAccountTransactionsUsecase(
  accountId: string,
  pagination?: PaginationDTO
): Promise<GetAccountTransactionsResponse> {
  const transactions = await prisma.transaction.findMany({
    orderBy: { occurredAt: "desc" },
    cursor: pagination?.cursor ? { id: pagination.cursor } : undefined,
    skip: pagination?.cursor ? 1 : 0,
    take: pagination?.limit,
    where: { accountId },
  });

  return transactions.map((transaction) => ({
    id: transaction.id,
    amount: transaction.amount.toNumber(),
    category: transaction.category,
    isExpense: transaction.isExpense,
    occurredAt: transaction.occurredAt,
  }));
}
