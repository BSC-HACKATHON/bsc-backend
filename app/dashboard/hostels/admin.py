from django.contrib import admin
from .models import Hauberge,Offer
from unfold.admin import ModelAdmin

@admin.register(Offer)
class OfferAdmin(ModelAdmin):
    list_display = ('id', 'name', 'description')  # Display relevant fields in the list view
    search_fields = ('name',) 


@admin.register(Hauberge)
class HaubergeAdmin(ModelAdmin):
    list_display = ('id', 'name', 'type', 'capacity', 'location', 'address', 'phone', 'reserved_people_count', 'availability')
    search_fields = ('name', 'location', 'address', 'phone')  # Enable search by key fields
    list_filter = ('type', 'availability')  # Add filters for type and availability
    autocomplete_fields = ('offers',)  # Use autocomplete for Many-to-Many relationships
    readonly_fields = ('reserved_people_count',)  # Make `reserved_people_count` read-only



# Register your models here.
