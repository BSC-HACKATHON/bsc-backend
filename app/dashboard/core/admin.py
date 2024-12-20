from django.contrib import admin
from django.contrib.auth.admin import UserAdmin as BaseUserAdmin
from .models import User
from unfold.admin import ModelAdmin
@admin.register(User)
class UserAdmin(BaseUserAdmin,ModelAdmin):
    # Fields to display in the list view
    list_display = (
        'email', 
        'first_name', 
        'last_name', 
        'identity_card_number', 
        'is_staff', 
        'is_active'
    )
    
    # Fields to use for search functionality
    search_fields = (
        'email', 
        'first_name', 
        'last_name', 
        'identity_card_number'
    )
    
    # Filters for the list view
    list_filter = (
        'is_staff', 
        'is_active', 
        'gender'
    )
    
    # Fields to display on the user edit form
    fieldsets = (
        (None, {
            'fields': ('email', 'password')
        }),
        ('Personal Info', {
            'fields': ('first_name', 'last_name', 'identity_card_number', 'date_of_birth', 'place_of_birth', 'gender')
        }),
        ('Permissions', {
            'fields': ('is_staff', 'is_active', 'is_superuser', 'groups', 'user_permissions')
        }),
        ('Important Dates', {
            'fields': ('last_login',)
        }),
    )
    
    # Fields to display when creating a new user
    add_fieldsets = (
        (None, {
            'classes': ('wide',),
            'fields': (
                'email', 
                'password1', 
                'password2', 
                'first_name', 
                'last_name', 
                'identity_card_number', 
                'date_of_birth', 
                'place_of_birth', 
                'gender', 
                'is_staff', 
                'is_active'
            ),
        }),
    )
    
    # Fields for ordering in the list view
    ordering = ('email',)
    
    # Sets the read-only fields
    readonly_fields = ('last_login',)

