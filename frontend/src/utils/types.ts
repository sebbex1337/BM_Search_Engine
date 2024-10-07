export type Page = {
  title: string;
  URL: string;
  Language: string;
  LastUpdated: string;
  Content: string;
};

export type AuthResponse = {
  statusCode: number;
  message: string;
  username?: string; // Does not need to be there
};
