import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, throwError } from 'rxjs';
import { environment } from '../../environments/environment';

export interface Hello {
  id: number,
  text: string,
  createdAt: string,
}
@Injectable({
  providedIn: 'root'
})
export class HelloService {


  constructor(private http: HttpClient) { }

  getHellos(): Observable<Hello> {
    return this.http.get<Hello>(this.basePath() + "/hello").pipe(catchError(this.handleError))
  }

  basePath(): string {
    return environment.basePath
  }

  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      console.error('An error occurred:', error.error);
    } else {
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }
}
