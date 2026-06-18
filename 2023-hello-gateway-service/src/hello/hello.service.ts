import { HttpService, Injectable } from '@nestjs/common';
import { map } from 'rxjs/operators';

@Injectable()
export class HelloService {
  constructor(private httpService: HttpService) {}

  hello() {
    const externalHelloService = process.env.HELLO_SERVICE || 'http://localhost:3000/hello';
    console.log('ATTEMPTING: proxy on to', externalHelloService);
    return this.httpService.get(externalHelloService).pipe(map(response => response.data));
  }
}
