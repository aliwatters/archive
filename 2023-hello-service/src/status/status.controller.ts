import { Controller, Get } from '@nestjs/common';

@Controller('status')
export class StatusController {
  @Get()
  status(): object {
    return { status: "OK"};
  }
}
