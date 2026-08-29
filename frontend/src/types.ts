export interface Entry {
    id: number;
    record_date: string;
    record_title: string;
    record_type: string;
}

export type NewEntry = {
  record_date: string
  record_title: string
  record_type: string
}