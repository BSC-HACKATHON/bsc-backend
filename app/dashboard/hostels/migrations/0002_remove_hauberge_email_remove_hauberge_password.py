# Generated by Django 4.2.17 on 2024-12-20 03:31

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('hostels', '0001_initial'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='hauberge',
            name='email',
        ),
        migrations.RemoveField(
            model_name='hauberge',
            name='password',
        ),
    ]
