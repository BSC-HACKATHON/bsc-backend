from django.db import models

class BlackList(models.Model):
    first_name = models.CharField(max_length=255)
    last_name = models.CharField(max_length=255)
    identity_card_number = models.CharField(max_length=50, unique=True)

    def __str__(self):
        return f"Blacklisted: {self.first_name} {self.last_name}"

