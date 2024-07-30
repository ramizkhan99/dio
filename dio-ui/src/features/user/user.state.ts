export interface UserState {
  /**
   * First name of the user
   */
  firstName: string;

  /**
   * Last name of the user
   */
  lastName: string;

  /**
   * Usename
   */
  username: string;

  /**
   * EmailID
   */
  emailId: string;

  /**
   * Is api call pending
   */
  isPending: boolean;

  hasError: boolean;
}

export const defaultState: UserState = {
  firstName: '',
  lastName: '',
  username: '',
  emailId: '',
  isPending: false,
  hasError: false
};