import { Controller, Get } from '@nestjs/common';
import { HelloService } from './hello.service';

@Controller('hello-proxy')
export class HelloController {
  constructor(private readonly helloService: HelloService) {}

  @Get()
  proxyHello() {
    return this.helloService.hello();
  }
}
