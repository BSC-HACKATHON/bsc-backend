# Generated by Django 4.2.17 on 2024-12-20 02:05

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Offer',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=255)),
                ('description', models.TextField()),
            ],
        ),
        migrations.CreateModel(
            name='Hauberge',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('type', models.CharField(choices=[('house', 'House'), ('camp', 'Camp')], max_length=6)),
                ('capacity', models.IntegerField()),
                ('name', models.CharField(max_length=255)),
                ('location', models.CharField(max_length=255)),
                ('address', models.CharField(max_length=255)),
                ('email', models.EmailField(max_length=255)),
                ('password', models.CharField(max_length=255)),
                ('phone', models.CharField(max_length=20)),
                ('reserved_people_count', models.IntegerField(default=0)),
                ('availability', models.BooleanField(default=True)),
                ('image_list', models.JSONField()),
                ('offers', models.ManyToManyField(to='hostels.offer')),
            ],
        ),
    ]
