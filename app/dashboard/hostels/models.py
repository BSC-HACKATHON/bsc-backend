from django.db import models


class Offer(models.Model):
    name = models.CharField(max_length=255)
    description = models.TextField()
    def __str__(self):
        return self.name


class Hauberge(models.Model):
    TYPE_CHOICES = [
            ('house', 'House'),
            ('camp', 'Camp'),
        ]
            
    type = models.CharField(max_length=6, choices=TYPE_CHOICES)
    capacity = models.IntegerField()
    name = models.CharField(max_length=255)
    location = models.CharField(max_length=255)
    address = models.CharField(max_length=255)
    phone = models.CharField(max_length=20)
    reserved_people_count = models.IntegerField(default=0)
    availability = models.BooleanField(default=True)
    image_list = models.JSONField() 
    offers = models.ManyToManyField(Offer)

    def __str__(self):
        return self.name
    
