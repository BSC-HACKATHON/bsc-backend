from django.contrib import admin
from .models import Reservation
from unfold.admin import ModelAdmin



@admin.register(Reservation)
class ReservationAdmin(ModelAdmin):
    list_display = ('id', 'hauberge', 'resident', 'room_number', 'check_in_date', 'check_out_date', 'reservation_type', 'catering_amount')
    search_fields = ('hauberge__name', 'resident__username', 'room_number')  # Enable search by hostel name, resident username, and room number
    list_filter = ('reservation_type', 'check_in_date', 'check_out_date', 'hauberge')  # Add filters for common fields
    date_hierarchy = 'check_in_date'  # Add a date-based navigation bar
    autocomplete_fields = ['hauberge', 'resident']  # Enable autocomplete for related fields to improve usability
    ordering = ('-check_in_date',)  # Default ordering by most recent check-in
    readonly_fields = ('catering_amount',)  # Prevent editing catering amount (if needed)





# Register your models here.
