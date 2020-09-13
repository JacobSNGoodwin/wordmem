export interface Word {
  id: string; // optional as this will be created in postgres
  userId: string;
  email: string;
  word: string;
  refUrl: string;
  emailReminder: Boolean;
  startDate: Date;
}
