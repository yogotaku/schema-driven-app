export interface User {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  emailVerified: boolean;
  dateOfBirth: string;
  createDate?: string;
}

export const isUser = (arg: unknown): arg is User => {
  const u = arg as User;

  return (
    typeof u?.id === 'number' &&
    typeof u?.firstName === 'string' &&
    typeof u?.lastName === 'string' &&
    typeof u?.email === 'string' &&
    typeof u?.emailVerified === 'boolean' &&
    typeof u?.dateOfBirth === 'string' &&
    typeof u?.createDate === 'string'
  );
};

export interface NewUser {
  firstName: string;
  lastName: string;
  email: string;
  dateOfBirth?: string;
}

export interface Pet {
  id: number;
  tag: string;
  name: string;
}

export const isPet = (arg: unknown): arg is Pet => {
  const p = arg as Pet;

  return (
    typeof p?.id === 'number' &&
    typeof p?.tag === 'string' &&
    typeof p?.name === 'string'
  );
};
