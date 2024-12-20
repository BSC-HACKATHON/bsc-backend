from django.db import models
from hostels.models import Hauberge
from core.models import User

class Reservation(models.Model):
    RESERVATION_TYPE_CHOICES = [
        ('Free', 'Free'),
        ('Paid', 'Paid'),
    ]

    hauberge = models.ForeignKey(Hauberge, on_delete=models.CASCADE, related_name='reservations')
    resident = models.ForeignKey(User, on_delete=models.CASCADE, related_name='reservations')
    room_number = models.PositiveIntegerField()
    check_in_date = models.DateField()
    check_out_date = models.DateField()
    reservation_type = models.CharField(max_length=10, choices=RESERVATION_TYPE_CHOICES)
    catering_amount = models.DecimalField(max_digits=10, decimal_places=2)

    def __str__(self):
        return f"Reservation #{self.id} at {self.hauberge.name}"


