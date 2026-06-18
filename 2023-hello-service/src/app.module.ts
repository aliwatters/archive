import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { StatusController } from './status/status.controller';
import { HelloController } from './hello/hello.controller';

@Module({
  imports: [],
  controllers: [AppController, StatusController, HelloController],
  providers: [AppService],
})
export class AppModule {}
