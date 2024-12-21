from __future__ import absolute_import, unicode_literals
import os
import pika
import json
from celery import Celery
from django.contrib.auth import get_user_model
from celery.signals import worker_ready

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'dashboard.settings')

app = Celery('dashboard')

app.config_from_object('django.conf:settings', namespace='CELERY')

app.autodiscover_tasks()


@app.task(bind=True)
def debug_task(self):
    print(f'Request: {self.request!r}')

@app.task(bind=True)
def rabbitmq_consumer(self):
    # Connect to RabbitMQ
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()

    # Declare the queue
    channel.queue_declare(queue='users')

    # Callback function to process messages
    def callback(ch, method, properties, body):
        message = json.loads(body)
        print(message)
        user = get_user_model()
        user.objects.create_user(message["email"],message["first_name"],message["last_name"],message["identity_card_number"],password=message["password"],date_of_birth=message["date_of_birth"],place_of_birth=message["place_of_birth"],gender=message["gender"])
        print(f"Received message: {message}")

    # Start consuming messages
    channel.basic_consume(queue='users', on_message_callback=callback, auto_ack=True)

    print('Waiting for messages. To exit press CTRL+C')
    try:
        channel.start_consuming()
    except Exception as e:
        print(f"Error: {e}")
        connection.close()
        raise self.retry(exc=e, countdown=5)  # Retry in case of errors



@worker_ready.connect
def at_start(sender, **k):
    with sender.app.connection() as conn:
         sender.app.send_task('dashboard.celery.rabbitmq_consumer')
