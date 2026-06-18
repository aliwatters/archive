import { HttpModule, Module } from '@nestjs/common';
import { HelloController } from './hello.controller';
import { HelloService } from './hello.service';


@Module({
  imports: [HttpModule.register({
    timeout: 5000,
    maxRedirects: 5,
  })],
  controllers: [HelloController],
  providers: [HelloService],
})
export class HelloModule {}