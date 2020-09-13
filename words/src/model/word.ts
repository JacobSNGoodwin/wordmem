export interface Word {
  id: string; // optional as this will be created in postgres
  userId: string;
  email: string;
  word: string;
  definition: string;
  refUrl: string;
  emailReminder: Boolean;
  startDate: Date;
}

export const wordFromData = (dataObj: any): Word => ({
  id: dataObj.id,
  email: dataObj.email,
  emailReminder: dataObj.email_reminder,
  refUrl: dataObj.ref_url,
  startDate: dataObj.start_date,
  userId: dataObj.userid,
  word: dataObj.word,
  definition: dataObj.definition,
});
