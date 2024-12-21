from django.contrib.auth.models import AbstractBaseUser, PermissionsMixin,BaseUserManager
from django.db import models
from django.utils.translation import gettext_lazy as _

from django.contrib.auth.models import BaseUserManager

class UserManager(BaseUserManager):
    """
    Custom manager for the AbstractBaseUserModel to handle user creation.
    """
    def create_user(self, email, first_name, last_name, identity_card_number, password=None, **extra_fields):
        if not email:
            raise ValueError("The Email field must be set.")
        if not identity_card_number:
            raise ValueError("The Identity Card Number field must be set.")
        
        email = self.normalize_email(email)
        extra_fields.setdefault('is_active', True)
        user = self.model(
            email=email,
            first_name=first_name,
            last_name=last_name,
            identity_card_number=identity_card_number,
            **extra_fields
        )
        user.set_password(password)
        user.save(using=self._db)
        return user

    def create_superuser(self, email, first_name, last_name, identity_card_number, password=None, **extra_fields):
        extra_fields.setdefault('is_staff', True)
        extra_fields.setdefault('is_superuser', True)

        if extra_fields.get('is_staff') is not True:
            raise ValueError("Superuser must have is_staff=True.")
        if extra_fields.get('is_superuser') is not True:
            raise ValueError("Superuser must have is_superuser=True.")
        
        return self.create_user(email, first_name, last_name, identity_card_number, password, **extra_fields)


class User(AbstractBaseUser, PermissionsMixin):
    email = models.EmailField(unique=True)
    first_name = models.CharField(max_length=255)
    last_name = models.CharField(max_length=255)
    date_of_birth = models.DateTimeField()
    place_of_birth = models.CharField(max_length=255)
    gender = models.CharField(
        max_length=10,
        choices=[('Male', 'Male'), ('Female', 'Female')],
    )
    objects = UserManager()
    identity_card_number = models.CharField(max_length=50, unique=True)
    is_active = models.BooleanField(default=True)
    is_staff = models.BooleanField(default=False)
    USERNAME_FIELD = 'email'
    REQUIRED_FIELDS = ['first_name', 'last_name', 'identity_card_number', 'date_of_birth', 'place_of_birth']


    def __str__(self):
        return f"{self.first_name} {self.last_name} ({self.email})"

