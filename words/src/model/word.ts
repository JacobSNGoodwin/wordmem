export interface Word {
  id: string;
  userId: string;
  email: string;
  word: string;
  refUrl?: string;
  emailReminder: Boolean;
  startDate?: string;
}
