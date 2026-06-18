import { Module } from '@nestjs/common';
import { HelloModule } from './hello/hello.module';
import { StatusController } from './status/status.controller';

@Module({
  imports: [HelloModule],
  controllers: [StatusController],
  providers: [],
})
export class AppModule {}
