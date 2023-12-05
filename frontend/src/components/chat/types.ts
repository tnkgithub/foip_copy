export interface ChatData {
  user: string
  time: string
  table: string
  message: string
}

export const LengthMax = {
  USER: 100,
  TABLE: 100,
  MESSAGE: 500,
  TOTAL: 1024,
}
