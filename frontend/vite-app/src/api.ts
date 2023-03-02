import axios from 'axios';
import { isUser, type User, type NewUser, type Pet, isPet } from 'type';

const apiClient = axios.create({
  // baseURL: import.meta.env.VITE_PRISM_MOCK_URL,
  baseURL: import.meta.env.VITE_API_ENDPOINT_URL,
  timeout: 10000,
  paramsSerializer: { indexes: null },
  headers: { 'Content-Type': 'application/json' },
});

export const getUserById = async (id: number): Promise<User | null> => {
  const response = await apiClient({
    method: 'get',
    url: `/users/${id}`,
  });

  if (isUser(response.data)) {
    return response.data;
  }

  return null;
};

export const createUser = async (user: NewUser): Promise<User | null> => {
  const { firstName, lastName, email, dateOfBirth } = user;
  try {
    const response = await apiClient({
      method: 'post',
      url: '/users',
      data: {
        firstName,
        lastName,
        email,
        ...(dateOfBirth?.length !== 0 ? { dateOfBirth } : {}),
      },
    });

    if (isUser(response.data)) {
      return response.data;
    }

    return null;
  } catch {
    return null;
  }
};

export const patchUser = async (
  id: number,
  user: NewUser
): Promise<boolean> => {
  const { firstName, lastName, email, dateOfBirth } = user;
  try {
    const _ = await apiClient({
      method: 'patch',
      url: `/users/${id}`,
      data: {
        firstName,
        lastName,
        email,
        ...(dateOfBirth?.length !== 0 ? { dateOfBirth } : {}),
      },
    });

    return true;
  } catch {
    return false;
  }
};

export const getPets = async (
  tags: string[],
  limit: number | null
): Promise<Pet[] | null> => {
  const response = await apiClient({
    method: 'get',
    url: `/pets`,
    params: {
      ...(tags.length !== 0 ? { tags } : {}),
      ...(limit !== null ? { limit } : {}),
    },
  });

  if (Array.isArray(response.data) && response.data.every(isPet)) {
    return response.data;
  }

  return null;
};
