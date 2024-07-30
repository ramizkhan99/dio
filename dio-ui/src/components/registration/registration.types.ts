export interface RegistrationProps {
  /**
   * Username of the user 
   */
  username: string;

  /**
   * Email id of the user
   */
  emailId: string;

  /**
   * First name of the user
   */
  firstName: string;
  
  /**
   * Last name of the user
   */
  lastName: string;

  isPending: boolean;

  hasError: boolean;

  registerUser: (request: UserRegistrationRequest) => void;

  setFirstName: (firstName: string) => void;
}

export interface UserRegistrationRequest {
  /**
   * Username of the user 
   */
  username: string;

  /**
   * Email id of the user
   */
  emailId: string;

  /**
   * First name of the user
   */
  firstName: string;
  
  /**
   * Last name of the user
   */
  lastName: string;

  /**
   * Encrypted password of the user
   */
  password: string;
}

export interface UserRegistrationResponse {
  [key: string]: string;
}