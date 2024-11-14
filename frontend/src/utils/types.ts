export type Page = {
  title: string;
  url: string;
  Language: string;
  LastUpdated: string;
  Content: string;
};

export type AuthResponse = {
  statusCode: number;
  message: string;
  username?: string; // Does not need to be there
};

export interface weather {
  main: {
    temp: number;
  };
  name: string;
}
