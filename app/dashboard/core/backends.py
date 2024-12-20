from django.contrib.auth.backends import BaseBackend
from django.contrib.auth import get_user_model

class EmailBackend(BaseBackend):
    model = get_user_model()
    
    def authenticate(self,request=None,email=None,password=None):
        try:

            print("cc")
            user = self.model.get(email=email)
        except self.model.DoesNotExist:
            return None
        else:
            if user.check_password(password):
                return user
        return None


    def get_user(self,pk):
        try:
            user = self.model.get(pk=pk)
            return user
        except self.model.DoesNotExist:
            return None

