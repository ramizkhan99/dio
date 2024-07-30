import { httpClient } from '@api/http-common';
import { UserRegistrationRequest } from './registration.types';

export class RegistrationAPI {
  private baseUrl = '/users';
  private httpClient = httpClient;

  /**
   * Add new user by making a POST request to the API
   * @param request Request object containing new user details
   * @returns Axios response object containing either the successful response or the error
   */
  public async register(request: UserRegistrationRequest) {
    return await this.httpClient.post(
      this.baseUrl,
      request
    );
  }
}