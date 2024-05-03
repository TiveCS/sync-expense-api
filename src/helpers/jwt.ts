export function createTokenExpireTime(minute: number): number {
  return Math.floor(Date.now() / 1000) + 60 * minute;
}
