import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  const port = parseInt(process.env.SERVER_PORT || '3000', 10);
  await app.listen(port, () => console.log(`Microservice is listening on ${port}`));
}
bootstrap();
