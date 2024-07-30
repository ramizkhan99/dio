export interface IHttpClient {
  /**
   * Send a GET request to the endpoint
   * @param extension 
   */
  GET(url: string): object;

  /**
   * Send a POST request to the endpoint with payload
   * @param extension 
   * @param payload
   */
  POST<Req, Res>(url: string, payload: Req): Res;
}