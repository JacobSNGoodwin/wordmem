export interface Word {
  id: string;
  userId: string;
  word: string;
  definition: string;
  refUrl: string;
  emailReminder: Boolean;
  startDate: Date;
}
