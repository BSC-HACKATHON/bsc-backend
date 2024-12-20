from django.contrib import admin
from .models import BlackList
from unfold.admin import ModelAdmin

@admin.register(BlackList)
class BlackListAdmin(ModelAdmin):
    list_display = ('first_name','last_name','identity_card_number')
    search_fields = ('first_name','last_name','identiy_card_number')
# Register your models here.
